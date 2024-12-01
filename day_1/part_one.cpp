#include <cstdlib>
#include <fstream>
#include <string>
#include <iostream>
#include <vector>

#define INPUT_FILE "./input.txt"

// would be quicker to just sort both lists but this was more fun to program
int nthBiggest(int n, std::vector<int> vec){
  int *largest = new int[n];
  int max;
  int maxInd;
  for (int i = 0; i < n; i++){
    max = -1;
    maxInd = -1;
    for (int j = 0; j < vec.size(); j ++){
      if (vec[j] > max){
        max = vec[j];
        maxInd = j;
      }
    }
    vec.erase(vec.begin() + maxInd);
    largest[i] = max;
  }

  int ret = largest[n-1];
  delete[] largest;
  return ret;
}

int main(){
  std::vector<int> listA;
  std::vector<int> listB;
  std::ifstream inputfile(INPUT_FILE);
  std::string line;
  
  if (inputfile.is_open()){
    while (std::getline(inputfile, line)){
      if(line == ""){
        continue;
      }
      int spaceAt = line.find(" ");
      std::string a = line.substr(0, spaceAt);
      std::string b = line.substr(spaceAt + 1);
      listA.push_back(std::atoi(a.data()));
      listB.push_back(std::atoi(b.data())); 
    }
    inputfile.close();
  }else{
    std::cout << "Error while opening file!" << std::endl;
    return 1;
  }

  int distance = 0;
  int total = listA.size();
  for (int i = 1; i <= total; i++){
    distance += std::abs(nthBiggest(i, listA) - nthBiggest(i, listB));
  }

  std::cout << "Total distance: " << distance << std::endl;

  return 0;
}
