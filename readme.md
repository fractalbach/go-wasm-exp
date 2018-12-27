Experiments with webassembly and go
==================================================

Main Question: Which parts of an application should be included in the
webassembly, and which parts are best left outside?


In this experiment, only the portable logic will be in the
webassembly.  The exported functions can then be used.  The internal
game state is managed within web assembly, and then graphics and
interactions are left outside.
