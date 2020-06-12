(defsystem "woo"
  :depends-on ("woo"
               "clack-test"
               "rove")
  :components
  ((:file "t/woo")
   (:file "src/woo")
   (:file "t/ipv6"))
  :perform (test-op (op c) (symbol-call '#:rove '#:run c)))
