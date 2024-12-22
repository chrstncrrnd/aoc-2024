let break_number num =
  let num_str = string_of_int num in
  let len = String.length num_str in
  let half_len = len / 2 in
  let lhs = String.sub num_str 0 half_len in
  let rhs = String.sub num_str half_len half_len in
  [int_of_string lhs; int_of_string rhs]

let read_file filename =
  let file = open_in filename in
  input_line file

let extract_stones contents =
  let values = String.split_on_char ' ' contents in
  List.map int_of_string values

let num_digits number =
  let num_str = string_of_int number in
  String.length num_str


let blink_stone stone =
  if stone = 0 then
    [1]
  else if num_digits stone mod 2 = 0 then
    break_number stone
  else
    [stone * 2024]


let key_gen stone depth =
  (string_of_int stone ^ "," ^ string_of_int depth)

let () = 
  let stones = extract_stones (read_file "input.txt") in 
  let table = Hashtbl.create 10 in
  let rec traverse depth stone =
    if depth = 0 then
      1
    else
      if Hashtbl.mem table (key_gen stone depth) then
        Hashtbl.find table (key_gen stone depth)
      else
        let branches = blink_stone stone in
        let res = traverse (depth - 1) (List.nth branches 0) + if (List.length branches = 2)  then (traverse (depth - 1) (List.nth branches 1)) else 0 in
        Hashtbl.add table (key_gen stone depth) res;
        res 
  in
  let sum = List.fold_right ( + ) (List.map (traverse 75) stones) 0 in
  print_int sum;
