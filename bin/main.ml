open Views.Main
open Hello

let () = Dream.run (fun _ -> hello "Someone Else" |> main |> Dream.html)
