let main content = "<p>" ^ content ^ "</p>"

let%expect_test "main wraps in <p>" =
  main "Your content here" |> print_string;
  [%expect {| <p>Your content here</p> |}]
