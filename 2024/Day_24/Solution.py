from collections import defaultdict

# Only part 1

class Node:
  def __init__(self, label, operation, operands, value):
    self.label = label
    self.operation = operation
    self.operands = [operands[0], operands[1]]
    self.value = value

vals = {}
deps = defaultdict(Node)
with open("input.txt") as file:
  lines = file.read().split("\n")  
  input_mode_flag = True
  
  for ln in lines:
    if ln == "":
      input_mode_flag = False
      continue
    
    if input_mode_flag:
      bleh = ln.split(":")
      obj = Node(bleh[0], None, [None, None], int(bleh[1]))
      deps[bleh[0]] = obj
      vals[obj.label] = obj.value
    else:
      bleh = ln.split(" ")
      assert bleh[4] not in deps, "wire has multiple input lines"
      obj = Node(bleh[4], bleh[1], [bleh[0], bleh[2]], None)
      deps[bleh[4]] = obj

def parse(node):
  if node.value is not None:
    return node.value
  
  oper1 = node.operands[0]
  oper2 = node.operands[1]
  if oper1 in vals:
    op1 = vals[oper1]
  else:
    op1 = parse(deps[oper1])
  
  if oper2 in vals:
    op2 = vals[oper2]
  else:
    op2 = parse(deps[oper2])
  
  match node.operation:
    case "AND":
      node.value = op1 & op2
    case "OR":
      node.value = op1 | op2
    case "XOR":
      node.value = op1 ^ op2
    case default:
      assert 1 == 0, "Operation is not mapped"
  
  return node.value

for i in deps:
  vals[i] = parse(deps[i])

ans = 0
for i in vals:
  if i[0] == 'z' and vals[i] == 1:
    ans = ans ^ (1 << int(i[1:]))
print(ans) 