#include <fstream>
#include <iostream>
#include <unordered_map>

#define INPUT_FILE "./input.txt"

int main() {
  std::unordered_map<int, int> mapA;
  std::unordered_map<int, int> mapB;
  std::ifstream inputfile(INPUT_FILE);
  std::string line;
  
  if (inputfile.is_open()){
    while (std::getline(inputfile, line)){
      if(line == ""){
        continue;
      }
      int spaceAt = line.find(" ");
      int a = std::atoi(line.substr(0, spaceAt).data());
      int b = std::atoi(line.substr(spaceAt + 1).data());
      mapA[a] = mapA[a] + 1;
      mapB[b] = mapB[b] + 1;
    }
    inputfile.close();
  }else{
    std::cout << "Error while opening file!" << std::endl;
    return 1;
  }
  int similarity = 0;
  std::unordered_map<int, int>::iterator itr;
  std::pair<int, int> pair;
  for (itr = mapA.begin(); itr != mapA.end(); itr++){
    pair = *itr;
    if (mapB.find(pair.first) == mapB.end()){
      continue;
    }
    similarity += pair.first * pair.second * mapB[pair.first];
  }


  std::cout << "Similarity score: " << similarity << std::endl;

  return 0;

}
