let root c =
  {|<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width, initial-scale=1.0"><meta http-equiv="X-UA-Compatible" content="ie=edge"><title>Days</title></head><body>|}
  ^ String.concat "" c ^ {|</body></html>|}

let%expect_test "renders boilerplate around content" =
  [ "<p>test</p>"; "<p>other</>" ] |> root |> print_string;
  [%expect {| <!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width, initial-scale=1.0"><meta http-equiv="X-UA-Compatible" content="ie=edge"><title>Days</title></head><body><p>test</p><p>other</></body></html> |}]
