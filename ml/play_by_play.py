# NOTE: author Kendall Weihe

import tensorflow as tf
import numpy as np
import pdb
import os
import sys
import json

DATA = None
NUM_INPUTS = None
NUM_STEPS = 200
NUM_HIDDEN = 128
NUM_CLASSES = 1

def input_data(config):
    raw_data = []
    data_path = config["data_path"]
    files = os.listdir(data_path)
    for file in files:
        file_data = np.genfromtxt("{}/{}".format(data_path, file), delimiter=",")
        raw_data.append(file_data)
    DATA = np.array(raw_data)
    NUM_INPUTS = DATA.shape()[2]

def generate_model(config):
    x = tf.placeholder("float", [None, NUM_STEPS, NUM_INPUTS])
    y = tf.placeholder("float", [None])
    keep_prob = tf.placeholder("float")

    weights = {
        'out' : tf.get_variable("weights_1", shape=[NUM_HIDDEN, NUM_CLASSES],
                   initializer=tf.contrib.layers.xavier_initializer(), dtype=tf.float32),
    }
    biases = {
        'out': tf.Variable(tf.zeros([NUM_CLASSES]))
    }

    x = tf.transpose(x, [1, 0, 2])
    x = tf.reshape(x, [-1, NUM_INPUTS])
    x = tf.split(0, NUM_STEPS, x)
    lstm_cell = rnn_cell.BasicLSTMCell(NUM_HIDDEN, forget_bias=1.0)
    outputs, states = rnn.rnn(lstm_cell, x, dtype=tf.float32)
    output = tf.matmul(outputs[-1], weights['out']) + biases['out']
    output = tf.reshape(output, [-1])
    model_output =  tf.nn.dropout(output, keep_prob)

    return x, y, model_output

def train(config, x, y, model_output):
    n_samples = tf.cast(tf.shape(x)[0], tf.float32)
    cost = tf.reduce_sum(tf.pow(pred-y, 2))/(2*n_samples)
    optimizer = tf.train.AdamOptimizer(learning_rate=config["learning_rate"]).minimize(cost)
    accuracy = tf.reduce_mean(tf.abs(tf.sub(pred, y)))

    BATCH_SIZE = config["batch_size"]

    # Launch the graph
    with tf.Session() as sess:
        init = tf.global_variables_initializer()
        sess.run(init)

        # Shuffle data
        # TODO
        num_verification_games = DATA.shape[0] * config["verification_percentage"]
        training_data = DATA[0:DATA.shape[0]-num_verification_games, :, :]
        verification_data = DATA[DATA.shape[0]-num_verification_games:DATA.shape[0], :, :]

        # iterate through epochs
        for i in range(config["epochs"]):
            # iterate through training set
            for j in range(0, training_data.shape[0], BATCH_SIZE):
                batch_x = training_data[j:j+BATCH_SIZE, 0:NUM_STEPS, :]
                batch_y = training_data[j:j+BATCH_SIZE, -1, :]
                pdb.set_trace()
                print("reshape ground_truth")
                sess.run(optimizer, feed_dict={x: batch_x, y: batch_y, keep_prob: config["dropout"]})
                train_acc, train_loss = sess.run([accuracy, cost], feed_dict={x: batch_x, y: batch_y, keep_prob: 1.0})

                samples = sess.run(pred, feed_dict={x: verification_data, keep_prob: 1.0})
                pdb.set_trace()
                print("...")

def main():
    f = open("config.json", "r")
    config = json.loads(f.read())
    f.close()

    input_data(config)
    x, y, model_output = define_model(config)
    train(config, x, y model_output)

if __name__ == "__main__":
    main()
