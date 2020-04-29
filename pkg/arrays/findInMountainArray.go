package arrays

type MountainArray struct {
	Array []int
}

func (m *MountainArray) get(index int) int {
	return m.Array[index]
}
func (m *MountainArray) length() int {
	return len(m.Array)
}

// FindInMountainArray 山脉数组中查找目标值 https://leetcode-cn.com/problems/find-in-mountain-array
func FindInMountainArray(target int, mountainArr *MountainArray) int {
	left, right := 0, mountainArr.length()-1
	lens := right
	for left < right {
		mid := (left + right) / 2
		if mountainArr.get(mid) < mountainArr.get(mid+1) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	index := binarySearch(mountainArr, target, 0, left, 1)
	if index != -1 {
		return index
	}
	return binarySearch(mountainArr, target, left+1, lens, -1)
}

func binarySearch(mountainArr *MountainArray, target, left, right, flag int) int {
	target *= flag
	for left <= right {
		mid := (left + right) / 2
		cur := mountainArr.get(mid) * flag
		if cur == target {
			return mid
		} else if cur < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
