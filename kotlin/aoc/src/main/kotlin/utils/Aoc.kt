package utils

fun readFile(path: String) = object {}.javaClass.getResource(path)?.readText() ?: throw Exception("no file $path")

fun List<String>.listOfInts() = this.map { line -> line.split("\\s+".toRegex()).map { it.toInt() } }
