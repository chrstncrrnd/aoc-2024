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

let process_stones stones =
  print_endline ("List length: " ^ string_of_int (List.length stones));
  List.flatten (List.map blink_stone stones)

let rec process_n_times n stones iter =
  print_endline ("Blinking" ^ string_of_int iter);
  if iter >= n then
    stones
  else
    process_n_times n (process_stones stones) (iter + 1)


let () = 
  let stones = extract_stones (read_file "input.txt") in 
  print_endline "Part one total: ";
  let processed = process_n_times 25 stones 0 in
  print_int (List.length processed);
  print_newline ()

