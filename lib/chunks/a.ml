type target_options = Blank | Self

let a c ~link ~target =
  "<a href=\"" ^ link ^ "\" target=\""
  ^ (match target with Blank -> "_blank" | Self -> "_self")
  ^ "\">" ^ c ^ "</a>"

let%expect_test "a renders link" =
  a "test" ~link:"test.com" ~target:Blank |> print_string;
  [%expect {| <a href="test.com" target="_blank">test</a> |}]
