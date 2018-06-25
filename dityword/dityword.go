package dityword

import (
	"os"
	"log"
	"bufio"
	"io"
	"strings"
)

//256 tree
type dirtytree struct {
	bend bool
	subtree   [256]*dirtytree
}

var (
	dirtyhead *   dirtytree = nil
)

func loaddirtywords(filename string) bool {

	fi, err := os.Open(filename)
	if err != nil {
		log.Printf("filename=%v Error: %s\n",filename, err)
		return false
	}
	defer fi.Close()

	phead := new(dirtytree)

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
	//	log.Printf("a=%v \n",string(a))
		l := len(a)
		if l < 1 {
			continue
		}
		if l >256  {
			a = a[:256]
		}
		//fmt.Println(string(a))
		insertdirtywords(phead,a)
	}
	dirtyhead = phead
	return  true
}

func hasdirtywords(phead *dirtytree,str string) bool {
	if phead == nil {
		return  false
	}
	var pTree *dirtytree = phead
	//log.Printf("cmp string=%#v \n",str)
	strlower := []byte(strings.ToLower(string(str)))
	l := len([]byte(strlower))
	if l < 1  {
		return false
	}
	//log.Printf("cmp ToLower string=%#v \n",string(strlower))

	for i:=0;i< l;i++ {
		ch := byte(strlower[i])
		if pTree.subtree[ch] != nil {
			pTree = pTree.subtree[ch]
			if pTree.bend {
				return  true
			}
		}else {
			pTree = phead
		}
	}
	return false
}

//func filterdirtywords(phead *dirtytree,str string)  {
//
//}

func insertdirtywords(phead *dirtytree,str []byte)  {

	//全部小写
//	log.Printf("org        string=%#v \n",str)
	strlower := []byte(strings.ToLower(string(str)))
	l := len([]byte(strlower))
	if l < 1  {
		return
	}
//	log.Printf("org ToLower string=%#v \n",string(strlower))
//	log.Printf("org ToLower string=%#v \n",strlower)
	if phead  == nil {
		phead = new(dirtytree)
	}
	pTree := phead

	for i:=0;i< l;i++ {
		ch := byte(strlower[i])
		if pTree.subtree[ch]  == nil {
			pTree.subtree[ch] = new(dirtytree)
			pTree = pTree.subtree[ch]
		}
	}
	pTree.bend = true
}

//func releasedirtytree(phead *dirtytree)  {
	
//}


//api////////////////////////////////////////////////
func LoadDirtyWordsFile(filename string) bool  {
	return  loaddirtywords(filename)
}

func HasDirtyWords(chstr string) bool  {

	return hasdirtywords(dirtyhead,chstr)
}

//func FilterDirtyWords(filterstr string)  {
//
//}


