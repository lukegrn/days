let li c = String.concat "" @@ List.map (fun item -> "<li>" ^ item ^ "</li>") c
let ul c = "<ul>" ^ c ^ "</ul>"
let ol c = "<ol>" ^ c ^ "</ol>"

let lst c ~ordered =
  match ordered with true -> c |> li |> ol | false -> c |> li |> ul

let%expect_test "generates unordered list" =
  [ "1"; "2"; "3" ] |> lst ~ordered:false |> print_string;
  [%expect {| <ul><li>1</li><li>2</li><li>3</li></ul> |}]

let%expect_test "generates ordered list" =
  [ "1"; "2"; "3" ] |> lst ~ordered:false |> print_string;
  [%expect {| <ul><li>1</li><li>2</li><li>3</li></ul> |}]
