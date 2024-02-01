fun main() {
    val distances = mapOf(
        Pair("Tristram", "AlphaCentauri") to 34,
        Pair("Tristram", "Snowdin") to 100,
        Pair("Tristram", "Tambi") to 63,
        Pair("Tristram", "Faerun") to 108,
        Pair("Tristram", "Norrath") to 111,
        Pair("Tristram", "Straylight") to 89,
        Pair("Tristram", "Arbre") to 132,
        Pair("AlphaCentauri", "Snowdin") to 4,
        Pair("AlphaCentauri", "Tambi") to 79,
        Pair("AlphaCentauri", "Faerun") to 44,
        Pair("AlphaCentauri", "Norrath") to 147,
        Pair("AlphaCentauri", "Straylight") to 133,
        Pair("AlphaCentauri", "Arbre") to 74,
        Pair("Snowdin", "Tambi") to 105,
        Pair("Snowdin", "Faerun") to 95,
        Pair("Snowdin", "Norrath") to 48,
        Pair("Snowdin", "Straylight") to 88,
        Pair("Snowdin", "Arbre") to 7,
        Pair("Tambi", "Faerun") to 68,
        Pair("Tambi", "Norrath") to 134,
        Pair("Tambi", "Straylight") to 107,
        Pair("Tambi", "Arbre") to 40,
        Pair("Faerun", "Norrath") to 11,
        Pair("Faerun", "Straylight") to 66,
        Pair("Faerun", "Arbre") to 144,
        Pair("Norrath", "Straylight") to 115,
        Pair("Norrath", "Arbre") to 135,
        Pair("Straylight", "Arbre") to 127
    ).withDefault { 0 }

    val locations = distances.keys.flatMap { listOf(it.first, it.second) }.toSet()
    val shortestDistance = findShortestDistance(distances, locations)
    
    println("The shortest distance is: $shortestDistance")
}

fun findShortestDistance(distances: Map<Pair<String, String>, Int>, locations: Set<String>): Int {
    val allPermutations = locations.permutations()
    return allPermutations.minOf { permutation ->
        calculateTotalDistance(permutation, distances)
    }
}

fun calculateTotalDistance(route: List<String>, distances: Map<Pair<String, String>, Int>): Int {
    var totalDistance = 0
    for (i in 0 until route.size - 1) {
        val distance = distances.getValue(Pair(route[i], route[i + 1]))
        totalDistance += distance
    }
    return totalDistance
}

fun <T> Set<T>.permutations(): List<List<T>> {
    if (this.size <= 1) return listOf(this.toList())
    val perms = mutableListOf<List<T>>()
    val sub = this.first()
    for (perm in this.drop(1).toSet().permutations()) {
        for (i in 0..perm.size) {
            val newPerm = perm.toMutableList()
            newPerm.add(i, sub)
            perms.add(newPerm)
        }
    }
    return perms
}
