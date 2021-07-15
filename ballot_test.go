package main

import "testing"

func TestFindWinner(t *testing.T)  {
	testCases:=[]struct{
		candidateName []string
		winner string
	}{
		{[]string{"a"},"a"},
		{[]string{"a","b","b","a","A","c","C","C"},"c"},
		{[]string{},""},
	}

	for _,testCase:=range testCases{
		actualWinner:=FindWinner(testCase.candidateName)
		if actualWinner!=testCase.winner{
			t.Errorf("FindWinner didn't return the correct winner - expected %v, got %v",testCase.winner,actualWinner)
		}
	}
}

func TestFindWinnerV2(t *testing.T)  {
	testData:=make([]Ballot,0,0)
	testData=append(testData,Ballot{[]string{"c","a","a"}})
	testData=append(testData,Ballot{[]string{"c","a","a"}})
	testCases:=[]struct{
		candidateName []Ballot
		winner string
	}{
		{testData,"a"},
	}

	for _,testCase:=range testCases{
		actualWinner:=FindWinnerV2(testCase.candidateName,false)
		if actualWinner!=testCase.winner{
			t.Errorf("FindWinner didn't return the correct winner - expected %v, got %v",testCase.winner,actualWinner)
		}
	}
}
