function reverse(arr) {
    if (!arr.length) return null

    let i = 0;
    let j = arr.length - 1

    while (i < j) {
        const temp = arr[i];
        arr[i] = arr[j];
        arr[j] = temp;
        i +=  1
        j -= 1
        console.log(arr)
    }

    return arr
}


console.log(reverse([234, 56, 2, 4, 767, 3, 45, 23, 8]))