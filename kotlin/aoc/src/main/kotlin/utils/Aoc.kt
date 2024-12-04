package utils

fun readFile(path: String) = object {}.javaClass.getResource(path)?.readText() ?: throw Exception("no file $path")

fun List<String>.listOfInts() = this.map { line -> line.split("\\s+".toRegex()).map { it.toInt() } }

fun String.findAll(pattern: String) = Regex(pattern).findAll(this)

fun <T, R> Iterable<T>.mapIgnoreErrors(transform: (T) -> R): List<R> {
    val acc = mutableListOf<R>()
    for (item in this) {
        try {
            acc += transform(item)
        } catch (_: Exception) {
        }
    }
    return acc
}
