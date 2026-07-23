//// [tests/cases/compiler/tluaNoParameterProperties.tlua] ////

//// [tluaNoParameterProperties.tlua]
// Constructor parameter properties are removed with classes. Modifiers on an
// ordinary function parameter are always invalid.
function rejected(
  public a: number,
  private b: number,
  protected c: number,
  readonly d: number,
  override e: number,
): number
  return a + b + c + d + e;
end


//// [tluaNoParameterProperties.lua]
-- Constructor parameter properties are removed with classes. Modifiers on an
-- ordinary function parameter are always invalid.
function rejected(a, b, c, d, e)
    return a + b + c + d + e;
end
