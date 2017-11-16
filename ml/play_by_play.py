# NOTE: author Kendall Weihe

import tensorflow as tf
from tensorflow.python.ops import rnn_cell
import numpy as np
import pdb
import os
import sys
import json

DATA = None
GROUND_TRUTH = []
NUM_INPUTS = None
NUM_STEPS = 200
NUM_HIDDEN = 128
NUM_CLASSES = 1

def input_data(config):
    print("Reading data...")
    global NUM_INPUTS, GROUND_TRUTH, DATA
    raw_data = []
    raw_ground_truth = []
    data_path = config["data_path"]
    files = os.listdir(data_path)
    for file in files:
        file_data = np.genfromtxt("{}/{}".format(data_path, file), delimiter=",")
        raw_data.append(file_data[0:NUM_STEPS, :])
        raw_ground_truth.append(file_data[-1, 0])
        if not NUM_INPUTS:
            NUM_INPUTS = file_data.shape[1]
    DATA = np.array(raw_data)
    GROUND_TRUTH = np.array(raw_ground_truth)

def define_model(config, x, y, keep_prob, weights, biases):
    print("Generating model...")

    x = tf.transpose(x, [1, 0, 2])
    x = tf.reshape(x, [-1, NUM_INPUTS])
    x = tf.split(x, NUM_STEPS, 0)
    lstm_cell = rnn_cell.BasicLSTMCell(NUM_HIDDEN, forget_bias=1.0)
    outputs, states = tf.nn.static_rnn(lstm_cell, x, dtype=tf.float32)
    output = tf.matmul(outputs[-1], weights['out']) + biases['out']
    output = tf.reshape(output, [-1])
    model_output =  tf.nn.dropout(output, keep_prob)

    return model_output

def train(config, x, y, keep_prob, model_output):
    print("Training...")
    n_samples = tf.cast(tf.shape(x)[0], tf.float32)
    cost = tf.reduce_sum(tf.pow(model_output-y, 2))/(2*n_samples)
    optimizer = tf.train.AdamOptimizer(learning_rate=config["learning_rate"]).minimize(cost)
    accuracy = tf.reduce_mean(tf.abs(tf.subtract(model_output, y)))

    BATCH_SIZE = config["batch_size"]

    # Launch the graph
    with tf.Session() as sess:
        init = tf.global_variables_initializer()
        sess.run(init)

        # Shuffle data
        # TODO
        num_verification_games = DATA.shape[0] * config["verification_percentage"]
        training_data = DATA[0:int(DATA.shape[0]-num_verification_games)]
        verification_data = DATA[int(DATA.shape[0]-num_verification_games):int(DATA.shape[0])]

        # iterate through epochs
        for i in range(config["epochs"]):
            # iterate through training set
            for j in range(0, training_data.shape[0]-BATCH_SIZE, BATCH_SIZE):
                batch_x = training_data[j:j+BATCH_SIZE, :, :]
                batch_y = GROUND_TRUTH[j:j+BATCH_SIZE]
                sess.run(optimizer, feed_dict={x: batch_x, y: batch_y, keep_prob: config["dropout"]})

            print("Epoch: {}".format(i))
            train_acc, train_loss = sess.run([accuracy, cost], feed_dict={x: batch_x, y: batch_y, keep_prob: 1.0})
            print("Training\tAcc: {}\tLoss: {}".format(train_acc, train_loss))

            sample_ground_truth = GROUND_TRUTH[int(DATA.shape[0]-num_verification_games):int(DATA.shape[0])]
            samples = sess.run(model_output, feed_dict={x: verification_data, keep_prob: 1.0})
            print("Samples: {}".format(samples))
            abs_diff = np.absolute(np.subtract(samples, sample_ground_truth))
            avg_diff = np.mean(abs_diff)
            print("AVG PRED SPREAD: {}".format(avg_diff))

def main():
    f = open("config.json", "r")
    config = json.loads(f.read())
    f.close()

    input_data(config)

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

    model_output = define_model(config, x, y, keep_prob, weights, biases)
    train(config, x, y, keep_prob, model_output)

if __name__ == "__main__":
    main()
