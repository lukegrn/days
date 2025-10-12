let root c =
  let styles =
    ":root { font-family: Times New Roman; font-size: 1.15rem; }body { width: \
     80%; margin: auto; margin-top: 4rem; }"
  in
  {|<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width, initial-scale=1.0"><meta http-equiv="X-UA-Compatible" content="ie=edge"><title>Days</title><style>|}
  ^ styles ^ {|</style></head><body>|} ^ String.concat "" c ^ {|</body></html>|}

let%expect_test "renders boilerplate around content" =
  [ "<p>test</p>"; "<p>other</>" ] |> root |> print_string;
  [%expect
    {| <!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width, initial-scale=1.0"><meta http-equiv="X-UA-Compatible" content="ie=edge"><title>Days</title><style>:root { font-family: Times New Roman; font-size: 1.15rem; }body { width: 80%; margin: auto; margin-top: 4rem; }</style></head><body><p>test</p><p>other</></body></html> |}]
