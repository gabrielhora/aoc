(ns aoc.y2023.day2
  (:require [clojure.java.io :as io]
            [clojure.string :as str]))

(def input (slurp (io/resource "aoc/y2023/day2.txt")))

(defn parse-game
  "in -> 1 red, 2 green, 6 blue
   out -> {\"red\" 1 \"green\" 2, \"blue\" 6}"
  [text]
  (->> text
       str/trim
       (re-seq #"(\d+)\s(red|green|blue)")
       (map (fn [[_ num color]]
              [color (Integer/parseInt num)]))
       (into {})))

(defn parse-games
  "in -> 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
   out -> ({blue:3,red:4}, {red:1,green:2,blue:6}, {green:2})"
  [text]
  (map parse-game (str/split text #";")))

(defn parse-line
  "in -> Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
   out -> [1, [{blue:3,red:4}, {red:1,green:2,blue:6}, {green:2}]]"
  [input]
  (let [[id-part plays-part] (str/split input #":")
        id (re-find #"\d+" id-part)
        games (parse-games plays-part)]
    [(Integer/parseInt id) games]))

(defn parse
  [input]
  (->> input
       str/split-lines
       (map parse-line)
       merge))

(defn valid-game?
  [game]
  (every?
    (fn [play]
      (and (<= (get play "red" 0) 12)
           (<= (get play "green" 0) 13)
           (<= (get play "blue" 0) 14)))
    game))

(defn part1
  [input]
  (->> input
       parse
       (map (fn [[id game]]
              (if (valid-game? game) id 0)))
       (reduce +)))

(defn power-of
  [game]
  (let [reds (map #(get %1 "red" 0) game)
        greens (map #(get %1 "green" 0) game)
        blues (map #(get %1 "blue" 0) game)]
    (* (apply max reds)
       (apply max greens)
       (apply max blues))))

(defn part2
  [input]
  (->> input
       parse
       (map (fn [[_ game]] (power-of game)))
       (reduce +)))

(comment
  (part1 input)
  (part2 input)
  )