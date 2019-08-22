package kirk

/*
  #include "kirk_engine.h"
  #include "amctrl.h"
*/
import "C"

import "fmt"

func KirkInit() {
	fmt.Println(C.kirk_init())
}
