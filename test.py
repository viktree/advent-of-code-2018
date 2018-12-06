polymer = "dabAcCaCBAcCcaDA"

new_polymer = polymer
for i,unit  in enumerate(polymer):
    if(i + 1 != len(polymer)):
        before =    polymer[:i+1]
        after =     polymer[i+1:]
        next_unit = polymer[i+1]
        if(unit xor next_unit == 32):
            new_unit = ""
        else:
            new_unit = unit + next_unit
        print(i, before, unit, next_unit, after)