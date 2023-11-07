(ns aoc.y2020.day7
  (:require [aoc.y2020.utils :refer [split]]
            [clojure.string :as str]
            [clojure.walk :as walk]))

(def example "light red bags contain 1 bright white bag, 2 muted yellow bags.\ndark orange bags contain 3 bright white bags, 4 muted yellow bags.\nbright white bags contain 1 shiny gold bag.\nmuted yellow bags contain 2 shiny gold bags, 9 faded blue bags.\nshiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.\ndark olive bags contain 3 faded blue bags, 4 dotted black bags.\nvibrant plum bags contain 5 faded blue bags, 6 dotted black bags.\nfaded blue bags contain no other bags.\ndotted black bags contain no other bags.")

(defn str->keyword [s] (keyword (str/replace s " " "-")))

(defn parse-bag
  "Parse a bag and it's contents.

  Transforms
    light red bags contain 1 bright white bag, 2 muted yellow bags.
  Into
    {:light-red ([1 :bright-white [2 :muted-yellow])}
  "
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

(defn find-bag
  ([bags lookup-bag]
   (find-bag bags lookup-bag (first (keys bags))))

  ([bags lookup-bag start]
   (let [inside-bags (map second (get bags start))]
     (if (contains? inside-bags lookup-bag)
       true
       (find-bag))
     inside-bags)))

(comment
  (let [bags (parse example)]
    (find-bag bags :shiny-gold)
    )

  (parse-bag "light red bags contain 1 bright white bag, 2 muted yellow bags.")
  (parse-bag "bright white bags contain 1 shiny gold bag.")
  (parse-bag "faded blue bags contain no other bags.")
  )
