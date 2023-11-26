(ns aoc.y2020.day7
  (:require [aoc.y2020.utils :refer [split]]
            [clojure.java.io :as io]
            [clojure.string :as str]))

(def input (slurp (io/resource "aoc/y2020/day7.txt")))

(defn str->keyword [s] (keyword (str/replace s " " "-")))

(defn parse-bag
  "in:  light red bags contain 1 bright white bag, 2 muted yellow bags.
   out: {:light-red ([1 :bright-white [2 :muted-yellow])}"
  [line]
  (let [parts (re-seq #"(^|\d+)\s*(\w+ \w+) bag" line)
        [[_ _ bag] & contents] parts
        parsed-contents (for [[_ num name] contents]
                          [(Integer/parseInt num) (str->keyword name)])]
    {(str->keyword bag) parsed-contents}))

(defn parse [input]
  (->> input
       (split #"\n")
       (map parse-bag)
       (into {})))

(defn contains-bag?
  [all-bags bag lookup]
  (->> bag
       all-bags
       (some (fn [[_ colour]]
               (or
                 (= colour lookup)
                 (contains-bag? all-bags colour lookup))))))

(defn part1 [input]
  (let [all-bags (parse input)]
    (->> all-bags
         (filter (fn [[bag _]]
                   (contains-bag? all-bags bag :shiny-gold)))
         count)))

(defn count-bags
  [all-bags lookup]
  (->> lookup
       all-bags
       (reduce
         (fn [acc [n bag]]
           (+ acc n (* n (count-bags all-bags bag))))
         0)))

(defn part2 [input]
  (let [all-bags (parse input)]
    (count-bags all-bags :shiny-gold)))

(comment
  (part1 input)
  (part2 input)
  )

