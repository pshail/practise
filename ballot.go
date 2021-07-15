package main

import (
	"fmt"
	"strings"
)

type Ballot struct {
	//Weighted votings
	Votes []string
}

//FindWinner V2 - Input - List<Ballot>
//               Output - String
func FindWinnerV2(candidateBallots []Ballot,firstTieWin bool) string {
	ballotCount:=make(map[string]int)
	winner:=""
	maxVotes:=0
	for _,v:=range candidateBallots {
		for index,v:=range v.Votes{
			voteWeight:=3-index
			candidateName:=strings.ToLower(v)
			if _,OK:=ballotCount[candidateName];OK {
				ballotCount[candidateName]=ballotCount[candidateName]+voteWeight
			}else {
				ballotCount[candidateName]=voteWeight
			}
			if ballotCount[candidateName]>maxVotes{
				maxVotes=ballotCount[candidateName]
				winner=candidateName
				fmt.Println(fmt.Sprintf("Weight is %v for candidate %v",maxVotes,winner))
			}else if ballotCount[candidateName]==maxVotes{
				//what to do when val ==maxVotes
				if !firstTieWin {
					winner=candidateName
				}
			}
		}
	}
	return winner
}
// FindWinner  - Input - List<String> (Candidates' Name) - NFRs - 100k
//               Output - String
func FindWinner(candidateNames []string) string {
	ballotCount:=make(map[string]int)
	winner:=""
	maxVotes:=0
	for _,v:=range candidateNames {
		candidateName:=strings.ToLower(v)
		if _,OK:=ballotCount[candidateName];OK {
			ballotCount[candidateName]=ballotCount[candidateName]+1
		}else {
			ballotCount[candidateName]=1
		}
		if ballotCount[candidateName]>maxVotes{
			maxVotes=ballotCount[candidateName]
			winner=candidateName
		}else if ballotCount[candidateName]==maxVotes{
			//what to do when val ==maxVotes
			winner=candidateName
		}
	}
	return winner
}

func main(){
	candidateNames:=[]string{"a","b","b","a","c","C"}
	findWin:=FindWinner(candidateNames)
	fmt.Println(fmt.Sprintf("Winner is %v",findWin))
}
