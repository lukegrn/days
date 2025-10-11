let p c = "<p>" ^ c ^ "</p>"

let%expect_test "p wraps content in <p>" =
  p "content" |> print_string;
  [%expect {| <p>content</p> |}]
