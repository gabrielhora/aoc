(ns aoc.y2020.day1
  (:require [clojure.string :as str]))

(def example "1721\n979\n366\n299\n675\n1456")

(def input "1619\n1919\n1441\n1861\n1932\n1514\n1847\n1871\n1764\n1467\n1970\n1589\n2009\n1429\n1098\n1327\n1502\n1398\n1710\n1562\n1512\n1468\n1762\n1348\n1356\n1950\n1266\n1969\n1815\n1583\n1959\n1092\n1694\n1814\n1763\n1151\n1981\n1193\n1614\n1413\n1642\n1943\n1407\n895\n1430\n1706\n1962\n1522\n1486\n1986\n1623\n1489\n1411\n1851\n1817\n1416\n1654\n1438\n1419\n1649\n1362\n690\n1804\n1452\n1766\n1360\n1807\n1385\n1964\n1626\n1832\n745\n1702\n1602\n1471\n1996\n1915\n1813\n1460\n1925\n1638\n1581\n1584\n1379\n1148\n1554\n1564\n1914\n1757\n1820\n1559\n1096\n1944\n1587\n1499\n390\n1733\n1371\n1781\n2002\n324\n1655\n1639\n1482\n1198\n1264\n1953\n1320\n1704\n1321\n1449\n1455\n1509\n1765\n1797\n1703\n1758\n1610\n1756\n1901\n1707\n1968\n1601\n1328\n1336\n1592\n1678\n1699\n1793\n1957\n2000\n1306\n1094\n1545\n1331\n1751\n1739\n1335\n1753\n1983\n1966\n1934\n1831\n1426\n1711\n1840\n1857\n1347\n1789\n1409\n1310\n1752\n1897\n1497\n1485\n1125\n1803\n1577\n919\n1635\n1791\n1456\n1796\n1974\n1954\n1828\n2004\n1890\n1376\n1569\n1406\n1463\n2006\n1109\n1620\n1656\n1870\n1498\n1645\n1145\n1681\n1269\n1527\n1621\n1575\n1324\n1647\n1519\n1697\n1421\n1216\n1846\n1625\n1585\n1369\n1882\n1823\n1388\n1548\n1879\n")

(defn parse-int-lines [input]
  (let [lines (str/split-lines input)]
    (map #(Integer/parseInt %) lines)))

(defn part1 [input]
  (let [data (parse-int-lines input)
        sums (for [x data y data
                   :let [sum (+ x y)]
                   :when (= sum 2020)]
               [x y])
        result (first sums)]
    (reduce * result)))

(defn part2 [input]
  (let [data (parse-int-lines input)
        sums (for [x data y data z data
                   :let [sum (+ x y z)]
                   :when (= sum 2020)]
               [x y z])
        result (first sums)]
    (reduce * result)))

(comment
  (part1 input)
  (part2 input)
  )
