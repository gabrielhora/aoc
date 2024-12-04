package utils

fun readFile(path: String) = object {}.javaClass.getResource(path)?.readText() ?: throw Exception("no file $path")

fun List<String>.listOfInts() = this.map { line -> line.split("\\s+".toRegex()).map { it.toInt() } }

fun String.findAll(pattern: String) = Regex(pattern).findAll(this)

fun String.replaceCharAt(index: Int, newChar: Char): String {
    if (index < 0 || index >= length) {
        throw IndexOutOfBoundsException()
    }
    return substring(0, index) + newChar + substring(index + 1)
}