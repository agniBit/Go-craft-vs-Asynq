import glob

m = {}
keyQueue = {}
for file in glob.glob("./logs.txt"):
    with open(file, "r") as file:
        for line in file:
            try:
                l = line.split("]")[-1].strip()
                lst = l.split('"')
                if lst[1] not in m:
                    m[lst[1]] = 0
                m[lst[1]] += 1
                key = lst[1] + "----" + lst[3]
                if key not in keyQueue:
                    keyQueue[key] = 0
                keyQueue[key] += 1
            except:
                continue
list = sorted(m.items(), key=lambda x: x[1], reverse=True)
for (key, val) in list:
    print("{:<15} {:<20}".format(key, val))

list = sorted(keyQueue.items(), key=lambda x: x[1], reverse=True)

k = 0
for (key, val) in list:
    if k > 100:
        break
    print("{:<15} {:<75} {:<10}".format(key.split("----")[0], key.split("----")[1], val))
    k += 1
