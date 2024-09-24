## Garbage collector three color marking. 
Object has references of objects and colour. Initialy root object starts being gray, those objects on heap starts as white.

Black objects are reachable, gray objects are gonna be processed in the future and white objects are unreachable. 

Algorithm traverses through references of object, changing those objects' colours to gray. Once every object in the reference is gray, the main object becomes black. It continues to the moment there are no more gray objects to proceed. White objects that left on heap are gonna be collected by garbage collector.
