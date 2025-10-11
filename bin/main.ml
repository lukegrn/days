open Days

let () = Dream.run (fun _ -> Hello.hello "Someone Else" |> Dream.html)
