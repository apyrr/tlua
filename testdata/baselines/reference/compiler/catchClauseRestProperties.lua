//// [tests/cases/compiler/catchClauseRestProperties.tlua] ////

//// [catchClauseRestProperties.tlua]
try {
  // ...
} catch ({ ...rest }) {
  // ...
}


//// [catchClauseRestProperties.lua]
try {
    -- ...
}
catch ({}) { }
...;
rest;
{
    -- ...
}
