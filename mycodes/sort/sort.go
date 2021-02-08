package sort

func BubbleSort(s []int) {
	if len(s)<2 { return }
	for i:=0; i<len(s)-1; i++ {
		for j:=1; j<len(s)-i; j++ {
			if s[j-1]>s[j] {
				s[j-1], s[j] = s[j], s[j-1]
			}
		}
	}
}

func SelectionSort(s []int) {
	for i:=0; i<len(s); i++ {
		p := i
		for j:=i+1; j<len(s); j++ {
			if s[p]>s[j] {
				p=j
			}
		}
		s[i], s[p] = s[p], s[i]
	}
}

//直接插入排序
func InsertionSort(s []int) {
	for i:=1; i<len(s); i++ {
		v := s[i]
		p := i
		for j:=i-1; j>=0; j-- {
			if s[j]<=v { break }
			p = j
			s[j+1] = s[j]
		}
		s[p] = v
	}
}

//二分插入排序
func BinaryInsertionSort(s []int) {
	for i:=1; i<len(s); i++ {
		l, h := 0, i-1
		for l<=h {
			m := (l+h)/2
			if s[i]<s[m] {
				h = m-1
			} else {
				l = m+1
			}
		}
		v := s[i]
		for j:=i-1; j>=l; j-- {
			s[j], s[j+1] = s[j+1], s[j]
		}
		s[l] = v
	}
}

//增量插入排序
func ShellSort(s []int) {
	for a:=len(s)/2; a>0; a=a/2 {
		for i:=a; i<len(s); i++ {
			v := s[i]
			p := i
			for j:=i-a; j>=0; j-=a {
				if s[j]<=v { break }
				p = j
				s[j+a] = s[j]
			}
			s[p] = v
		}
	}
}

func merge(s []int, m int) {
	t := make([]int, m)
	copy(t, s[:m])
	i, j := 0, m
	for i<m&&j<len(s) {
		if t[i]>s[j] {
			s[i+(j-m)] = s[j]
			j++
		} else {
			s[i+(j-m)] = t[i]
			i++
		}
	}
	if i<m {
		copy(s[i+(j-m):], t[i:])
	}
}

func MergeSort(s []int) {
	if len(s)<2 { return }
	m := len(s)/2
	MergeSort(s[:m])
	MergeSort(s[m:])
	merge(s, m)
}

func partition(s []int) int {
	if len(s)<2 { return 0 }

	/*左右指针法*/
	v := s[0]
	i, j := 0, len(s)-1
	for i<j {
		/*以s[0]为基准，则必须先从右向左遍历*/
		for i<j&&s[j]>=v { j-- }
		for i<j&&s[i]<=v { i++ }
		s[i], s[j] = s[j], s[i]
	}
	s[0], s[i] = s[i], s[0]
	return i
	/*左右指针法*/

	// /*挖坑法*/
	// v := s[0]
	// i, j := 0, len(s)-1
	// for i<j {
	// 	for i<j&&s[i]<=v { i++ }
	// 	s[j] = s[i]
	// 	for i<j&&s[j]>=v { j-- }
	// 	s[i] = s[j]
	// }
	// s[i] = v
	// return i
	// /*挖坑法*/
}

func QuickSort(s []int) {
	p := partition(s)
	if p>0 { QuickSort(s[:p]) }
	if p<len(s) { QuickSort(s[p+1:]) }
}

func heapify(s []int, i int) {
	for j:=i; j<len(s); {
		p := 2*j+1
		if p+1<len(s)&&s[p]<s[p+1] {
			p = p+1
		}
		if p<len(s)&&s[j]<s[p] {
			s[j], s[p] = s[p], s[j]
			j = p;
		} else {
			break
		}
	}
}

func buildMaxHeap(s []int) {
	if len(s)<2 { return }
	for i:=(len(s)-1)/2; i>=0; i-- {
		heapify(s, i)
	}
}

func HeapSort(s []int) {
	buildMaxHeap(s)
	for i:=0; i<len(s); i++ {
		heapify(s[:len(s)-i], 0)
		s[0], s[len(s)-i-1] = s[len(s)-i-1], s[0]
	}
}

func CountingSort(s []int) {
	if len(s)<2 { return }
	min, max := s[0], s[0]
	for _,v:=range s {
		if min>v { min = v }
		if max<v { max = v }
	}
	b := make([]int, max-min+1)
	for _,v:=range s {
		b[v-min]++
	}
	p := 0
	for v,n:=range b {
		for i:=0; i<n; i++ {
			s[p] = v+min
			p++
		}
	}
}

//CountingSort是BucketSort的一个特例
func BucketSort(s []int) {
	if len(s)<2 { return }
	//计算最大值和最小值
	min, max := s[0], s[0]
	for _,v:=range s {
		if min>v { min = v }
		if max<v { max = v }
	}
	//计算桶的数量
	n := (max-min)/len(s)+1
	bs := make([][]int, n)
	//将元素放入桶中
	for _,v:=range s {
		bs[(v-min)/len(s)] = append(bs[(v-min)/len(s)], v)
	}
	//对每个桶排序
	for _,b:=range bs {
		InsertionSort(b)
	}
	//将桶中排好序的元素放回原序列
	l := 0
	for _,b:=range bs {
		copy(s[l:], b)
		l += len(b)
	}
}

func RadixSort(s []int) {
	max := s[0]
	for _,v:=range s {
		if max<v { max = v }
	}
	for r:=1; r<max; r*=10 {
		bs := make([][]int, 10)
		for _,v:=range s {
			bs[v/r%10] = append(bs[v/r%10], v)
		}
		l := 0
		for _,b:=range bs {
			copy(s[l:], b)
			l += len(b)
		}
	}
}