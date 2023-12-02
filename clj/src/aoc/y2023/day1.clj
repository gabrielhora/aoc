(ns aoc.y2023.day1
  (:require [clojure.java.io :as io]
            [clojure.string :as str]))

(def input (slurp (io/resource "aoc/y2023/day1.txt")))

(defn parse [input]
  (str/split-lines input))

(defn find-numbers-part1
  [line]
  (->> line
       seq
       (filter #(Character/isDigit ^Character %1))))

(def words->int {"zero"  0
                 "one"   1
                 "two"   2
                 "three" 3
                 "four"  4
                 "five"  5
                 "six"   6
                 "seven" 7
                 "eight" 8
                 "nine"  9})

(defn find-number-word
  [text]
  (->> (keys words->int)
       (filter #(str/ends-with? text %1))
       first
       (get words->int)))

(defn find-numbers-part2
  [text]
  (->> text
       seq
       (reduce
         (fn [[nums acc] char]
           (if (Character/isDigit ^Character char)          ; if is digit append numeric value to acc
             [(conj nums (Character/digit ^Character char 10)) acc]
             (let [new-acc (str acc char)                   ; otherwise try to find digit word and add
                   int-val (find-number-word new-acc)]      ; it's numeric value to the acc
               (if (nil? int-val)
                 [nums new-acc]
                 [(conj nums int-val) new-acc]))))
         [[] ""])
       first))

(defn sum-numbers
  [input search-fn]
  (->> input
       parse
       (map search-fn)
       (map (fn [nums] (Integer/parseInt (str (first nums) (last nums)))))
       (reduce +)))

(defn part1 [input]
  (sum-numbers input find-numbers-part1))

(defn part2 [input]
  (sum-numbers input find-numbers-part2))

(comment
  (part1 input)
  (part2 input)
  )