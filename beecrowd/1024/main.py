# solução aceita pelo avaliador do beecrowd

A = int(input())

for _ in range(A):
  line = list(input())

  for i in range(len(line)):
    if 'A' <= line[i] <= 'Z' or 'a' <= line[i] <= 'z':
      line[i] = chr(ord(line[i]) + 3)

  i, j = 0, len(line) - 1
  while i < j:
    line[i], line[j] = line[j], line[i]
    i, j = i + 1, j - 1

  for i in range(len(line) // 2, len(line)):
    line[i] = chr(ord(line[i]) - 1)

  print(''.join(line))
