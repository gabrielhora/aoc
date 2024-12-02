package utils

fun readFile(path: String) = object {}.javaClass.getResource(path)?.readText() ?: throw Exception("no file $path")

fun readLines(path: String) = readFile(path).lines()
