(ns aoc.y2023.day9
  (:require [clojure.java.io :as io]
            [clojure.string :as str]))

(def input (slurp (io/resource "aoc/y2023/day9.txt")))

(defn parse-line
  "convert string with numbers '1 2 3 4' to seq (1 2 3 4)"
  [line]
  (->> (str/split line #" ")
       (map #(Integer/parseInt %))))

(defn parse
  [input]
  (->> input
       str/split-lines
       (map parse-line)))

(defn next-layer
  "calculate the next layer by applying the diffs between numbers"
  [layer]
  (->> layer
       (partition 2 1)
       (map (fn [[a b]] (- b a)))))

(defn find-next
  [coll]
  (let [nl (next-layer coll)]
    (if (every? zero? nl)
      (last coll)
      (+ (last coll) (find-next nl)))))

(defn find-previous
  [coll]
  (let [nl (next-layer coll)]
    (if (every? zero? nl)
      (first coll)
      (- (first coll) (find-previous nl)))))

(defn part1
  [input]
  (->> input
       parse
       (map find-next)
       (reduce +)))

(defn part2
  [input]
  (->> input
       parse
       (map find-previous)
       (reduce +)))

(comment
  (part1 input)
  (part2 input))