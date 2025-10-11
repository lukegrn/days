(lang dune 3.20)

(name days)

(generate_opam_files true)

(source
 (github lukegrn/days))

(authors "Luke Green <luke@lukegreen.xyz>")

(maintainers "Luke Green <luke@lukegreen.xyz>")

(license MIT)

(documentation "")

(package
 (name days)
 (synopsis "A web application to host my \"A Photo a Day\" project")
 (description "")
 (depends
   ocaml
   dream
   ppx_expect
   ppx_inline_test)
 (tags
  (website "server side" photography images)))

; See the complete stanza docs at https://dune.readthedocs.io/en/stable/reference/dune-project/index.html
