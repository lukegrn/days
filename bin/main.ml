open Views.Main

let () = Dream.run @@ Dream.logger @@ fun _ -> main |> Dream.html
