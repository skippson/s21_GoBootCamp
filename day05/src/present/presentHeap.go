package present

type PresentHeap struct{
	root []Present
}

func (p PresentHeap) Len() int{
	return len(p.root)
}

func (p PresentHeap) less(i,j int) bool{
	status := false 
	if p.root[i].Value > p.root[j].Value{
		status = true
	} else if p.root[i].Value == p.root[j].Value {
		if p.root[i].Size < p.root[j].Size{
			status = true
		}
	}

	return status
}

func (p PresentHeap) Data() []Present{
	return p.root
}

func (p *PresentHeap) sort() {
	n := p.Len()
	for i := 0; i < n - 1; i++{
		idx := i
		for j := i + 1; j < n; j++{
			if p.less(j,idx){
				idx = j
			}
		}
		p.swap(i, idx)
	}
}

func (p PresentHeap) swap(i, j int) {
	p.root[i], p.root[j] = p.root[j], p.root[i]
}

func (p *PresentHeap) Push(new Present) {
	p.root = append(p.root, new)
	p.sort()
}

func (p *PresentHeap) Pop() Present{
	old := p.root
	val := old[0]
	p.root = old[1:]

	return val
}	

func HeapInit(p []Present) PresentHeap{
	root := &PresentHeap{}
	for _, val := range p{
		root.Push(val)
	}

	return *root
}