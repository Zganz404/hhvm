<?hh

function blah() :mixed{
  $xs = vec[11, 22, 33, 44, 55];
  $ys = vec['a', 'b', 'c', 'd', 'e', 'f'];
  $zs = vec['x', 'y', 'z'];

  try {
    echo "begin try\n";
    foreach ($zs as $z) {
      try {
        foreach ($ys as $y) {
          echo "begin outer loop $y\n";
          try {
            try {
              foreach ($xs as $x) {
                echo "begin inner loop $x\n";
                if ($x == 22) {
                  echo "continue\n";
                  continue;
                }
                echo "end inner loop $x\n";
              }
            } finally {
              echo "inner finally\n";
            }
          } finally {
            echo "outer finally\n";
          }
          echo "middle outer loop\n";
          try {
            $continue_outer_loop = false;
            foreach ($xs as $x) {
              try  {
                echo "begin inner loop $x\n";
                if ($x == 22 || $y == 'b') {
                  echo "continue 2\n";
                  $contine_outer_loop = true;
                  break;
                }
                if ($x == 33) {
                  echo "break\n";
                  break;
                }
                if ($y == 'f' || $z == 'y') {
                  echo "throw\n";
                  throw new Exception('ble!');
                }
              } finally {
                echo "exception finally\n";
              }
              echo "end inner loop $x\n";
            }
            if ($continue_outer_loop) continue;
          } finally {
            echo "inner finally 2\n";
          }
        }
        echo "end outer loop\n";
      } catch (Exception $e) {
        echo "catch\n";
      } finally {
        echo "outer finally 2\n";
      }
    }
  } finally {
    echo "final finally\n";
  }
}



<<__EntryPoint>>
function main_finally_break_007() :mixed{
blah();
}
