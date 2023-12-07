(ns aoc.y2023.day7
  (:require [clojure.java.io :as io]
            [clojure.string :as str]
            [clojure.test :refer [is]]))

(def input (slurp (io/resource "aoc/y2023/day7.txt")))

(defn in? [coll elm] (some #(= elm %) coll))

(def card-values {\2 2 \3 3 \4 4 \5 5 \6 6 \7 7 \8 8 \9 9 \T 10 \J 11 \Q 12 \K 13 \A 14})
(def card-values-joker {\J 2 \2 3 \3 4 \4 5 \5 6 \6 7 \7 8 \8 9 \9 10 \T 11 \Q 12 \K 13 \A 14})

(defn hand-value
  "convert a hand (string) to card values (seq of ints) based on a values table"
  [values-table card]
  (->> card
       seq
       (map #(get values-table %))))

(defn replace-jokers
  "find the best replacement for the joker"
  [hand]
  (let [hand-no-jokers (remove #(= 2 %) hand)
        sorted-freqs (->> hand-no-jokers
                          frequencies
                          (sort-by (juxt second first))     ; sort by number of occurences then value
                          reverse)
        replace-with (first (first sorted-freqs))]
    (replace {2 replace-with} hand)))

(comment
  (is (= '(7 3 14 3 3) (replace-jokers (hand-value card-values-joker "6JA22"))))
  (is (= '(11 6 6 6 6) (replace-jokers (hand-value card-values-joker "T55J5"))))
  (is (= '(13 11 11 11 11 (replace-jokers (hand-value card-values-joker "KTJJT")))))
  (is (= '(13 13 13 7 13) (replace-jokers (hand-value card-values-joker "KKK6J"))))
  )

(defn hand-weight
  "calculate the weight of the hand, five of a kind being the highest with 7, high card is 1"
  [hand]
  (let [freqs (->> hand replace-jokers frequencies vals)]
    (cond
      (in? freqs 5) 7
      (in? freqs 4) 6
      (and (in? freqs 3) (in? freqs 2)) 5
      (in? freqs 3) 4
      (= 2 (count (filter #(= 2 %) freqs))) 3
      (= 1 (count (filter #(= 2 %) freqs))) 2
      :else 1)))

(defn winning-hand
  "comparator for hands, returns -1 if 1 > 2, 0 if 1 == 2 and 1 if 1 < 2"
  [c1 c2]
  (let [w1 (hand-weight c1)
        w2 (hand-weight c2)]
    (if (= w1 w2)
      (loop [[h1 & t1] c1
             [h2 & t2] c2]
        (cond
          (nil? h1) 0                                       ; all cards are equal
          (> h1 h2) -1
          (< h1 h2) 1
          :else (recur t1 t2)))
      (compare w2 w1))))

(defn parse
  [values-table input]
  (->> input
       str/split-lines
       (map #(str/split % #"\s+"))
       (map (fn [[card bid]]
              (let [values (hand-value values-table card)]
                {:card   card
                 :values values
                 :type   (hand-weight values)
                 :bid    (Integer/parseInt bid)})))))

(defn winnings
  "calculate total winnings"
  [input values-table]
  (->> input
       (parse values-table)
       (sort-by :values winning-hand)
       reverse
       (map-indexed (fn [idx card] [(inc idx) card]))
       (reduce
         (fn [acc [rank {bid :bid}]]
           (+ acc (* rank bid)))
         0)))

(defn part1
  [input]
  (winnings input card-values))

(defn part2
  [input]
  (winnings input card-values-joker))

(comment
  (part1 input)
  (part2 input)
  )