package sort;

func quickSort(arr []int) {
    if len(arr) == 0 {
        return;
    }
}

func parSort(arr []int, l int, r int) {
    if l >= r {
        return;
    }
    m := partition(arr, l, r);
    parSort(arr, l, m - 1)
    parSort(arr, m + 1, r)
}

func partition(arr []int, l int, r int) int {
    j := l;
    target := arr[l];
    for i := l+1; i <= r; i++ {
        if arr[i] <= target {
            swap(arr[j+1], arr[i])
            j++
        }
    }
    swap(arr[j], arr[l])
    return j;
}

func swap(x int, y int) {
    t := y
    y = x
    x = t
}

