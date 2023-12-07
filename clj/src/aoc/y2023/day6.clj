(ns aoc.y2023.day6
  (:require [clojure.java.io :as io]
            [clojure.string :as str]))

(def input (slurp (io/resource "aoc/y2023/day6.txt")))

(defn parse-seq
  "return a sequence of vectors with [time distance]"
  [input]
  (let [[ts ds] (str/split-lines input)
        times (map #(Integer/parseInt %) (re-seq #"\d+" ts))
        distances (map #(Integer/parseInt %) (re-seq #"\d+" ds))]
    (map vector times distances)))

(defn parse-nums
  "return a vec of two numbers, time and distance"
  [input]
  (let [[ts ds] (str/split-lines input)]
    [(Long/parseLong (str/replace ts #"Time:(\s*)|\s*" ""))
     (Long/parseLong (str/replace ds #"Distance:(\s*)|\s*" ""))]))

(defn max-distance
  "each hold second increase speed by 1, return distance"
  [max-time hold-time]
  (let [rem-time (- max-time hold-time)]
    (* rem-time hold-time)))

(defn beat-record
  "return a list of times that can beat the record duration"
  [max-time record]
  (->> (+ 1 max-time)
       range
       (map
         (fn [time] [time (max-distance max-time time)]))
       (filter
         (fn [[_ dis]] (> dis record)))
       (map first)))

(defn part1
  [input]
  (->> input
       parse-seq
       (map #(apply beat-record %))
       (map count)
       (reduce *)))

(defn part2
  [input]
  (let [[time dist] (parse-nums input)]
    (count (beat-record time dist))))

(comment
  (part1 input)
  (part2 input))