let hello who = "Hello, " ^ who ^ "!"

let%expect_test "hello says hello to the name passed in" =
  hello "luke" |> print_string;
  [%expect {| Hello, luke! |}]
