#include <algorithm>
#include <bits/stdc++.h>

#define INPUT_FILE "./input.txt"


std::vector<int> splitString(const std::string str) {
    std::vector<int> result;
    std::string s = str;
    int splitAt = 0;
    while (splitAt != -1) {
      splitAt = s.find(",");
      result.push_back(std::atoi(s.substr(0, splitAt).data()));
      s.erase(0, splitAt + 1);
    }
    return result;
}

bool isValidUpdate(const std::vector<int> &update, const std::unordered_map<int, std::vector<int>> &rules){
  for (int i = 0; i < update.size(); ++i){
    int value = update[i];
    if (rules.count(value) == 0){
      continue;
    }
    std::vector<int> allowedPrev = rules.at(value);
    for (int j = i; j < update.size(); j++){
      if (std::count(allowedPrev.begin(), allowedPrev.end(), update[j]) > 0){
        return false;
      }
    }
  }
  return true;
}

bool comp(int a, int b, const std::unordered_map<int, std::vector<int>> &rules){
  if (rules.count(b) == 0){
    return false;
  }
  std::vector<int> prev = rules.at(b);
  return std::count(prev.begin(), prev.end(), a) > 0;
}

void sortUpdate(std::vector<int>* update, const std::unordered_map<int, std::vector<int>> &rules){
  std::sort(update->begin(), update->end(), [&](int a, int b){return comp(a,b, rules);});
}

int main() {
  std::ifstream inputfile(INPUT_FILE);
  std::string line;
  std::unordered_map<int, std::vector<int>> rules;
  std::vector<std::vector<int>> updates;
  int stage = 1;

  if (inputfile.is_open()){
    while (std::getline(inputfile, line)){
      if(line == ""){
        stage = 2; 
        continue;
      }
      if (stage == 1){
        int splitAt = line.find("|");
        int lhs = std::atoi(line.substr(0, splitAt).data());
        int rhs = std::atoi(line.substr(splitAt + 1).data());
        if (rules.count(rhs) == 0){
          rules.insert({rhs, {}});
        }
        rules[rhs].push_back(lhs);
      }else{
        updates.push_back(splitString(line));
      }
    }
    inputfile.close();
  }else{
    std::cout << "Error while opening file!" << std::endl;
    return 1;
  }

  int total = 0;
  for(std::vector update : updates){
    if (!isValidUpdate(update, rules)){
      sortUpdate(&update, rules);
      total += update[update.size()/2];
    }
  }

  std::cout << "Total is: " << total << std::endl;

  return 0;

}
