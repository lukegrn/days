let main =
  [
    Chunks.P.p
    @@ "Days is a photo project that will begin in 2026 and continue for the \
        full year in which I take at least one picture every day, and upload \
        one to this site. The purpose of this is two-fold: ";
    Chunks.Lst.lst
      [
        "Encourage me to shoot more";
        "Document daily life in its mundanity as well as its exceptionality";
      ]
      ~ordered:false;
    "Click " ^ Chunks.A.a "here" ~target:Self ~link:"/" ^ " to view the project";
  ]
  |> Chunks.Root.root
