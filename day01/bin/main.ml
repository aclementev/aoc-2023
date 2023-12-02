let read_lines filename = 
    let ic = open_in filename in
    let try_read () = 
        try Some (input_line ic) with End_of_file -> None in
    let rec loop acc = match try_read () with
        | Some s -> loop (s :: acc)
        | None -> close_in ic; List.rev acc in
    loop []


(* let is_alpha = function 'a' .. 'z' | 'A' .. 'Z' -> true | _ -> false *)

let is_digit = function '0' .. '9' -> true | _ -> false

(* Get the first element of a list *)
let first = function
    | [] -> failwith "empty list"
    | x :: _ -> x

(* Get the last element of a list *)
let rec last = function
    | [] -> failwith "empty list"
    | [x] -> x
    | _ :: rst -> last rst

(* Find the first digit in a string *)
let first_digit string = 
    string |>
    String.to_seq |> 
    Seq.filter is_digit |>
    Seq.take 1 |>
    List.of_seq |>
    first

(* Find the first digit in a string *)
let last_digit string = 
    string |>
    String.to_seq |> 
    Seq.filter is_digit |>
    List.of_seq |>
    last


let first_and_last_digit string = 
    let fst = first_digit string in
    let lst = last_digit string in
    (fst, lst)


(* Solution for the part 1 *)
let part_1 filename =
    let sum = read_lines filename |>
        List.map first_and_last_digit |>
        List.map (fun (x, y) -> Printf.sprintf "%c%c" x y) |>
        List.map int_of_string |>
        List.fold_left (+) 0
    in
    Printf.printf "%d" sum

(* let () = part_1 "sample.txt" *)
let () = part_1 "input.txt"
