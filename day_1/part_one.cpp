#include <algorithm>
#include <cstdlib>
#include <fstream>
#include <string>
#include <iostream>
#include <vector>

#define INPUT_FILE "./input.txt"

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
  sort(listA.begin(), listA.end(), [](int a, int b) {return a>b;});
  for (int i = 1; i <= total; i++){
    distance += std::abs(listA[i] - listB[i]);
  }

  std::cout << "Total distance: " << distance << std::endl;

  return 0;
}
