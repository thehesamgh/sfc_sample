# Service Function Chaining (SFC) sample in Golang
This is a simple way that you can create Service Function Chaining (SFC) using Chain of Responsibilitys design pattern in Golang.

## sample config file (YAML)
```yaml
sfc:
  chain: 
    - "firewall1"
    - "ids1"
  network_functions_info: 
    - name_of_function: "test_firewall_1"
      type_of_function: "firewall"
    - name_of_function: "test_ids_1"
      type_of_function: "ids"
    - name_of_function: "test_ids_2"
      type_of_function: "ids"

```

## sample output
```
Chain has been created: 
	test_firewall_1 --> test_ids_1 --> test_ids_2

test_firewall_1 is Processing Packet: "Hey I'm packet 1".
test_firewall_1 is Adding a Rule.
test_ids_1 is Processing Packet: "Hey I'm packet 1".
test_ids_2 is Processing Packet: "Hey I'm packet 1".
test_firewall_1 is Processing Packet: "Hey I'm packet 2".
test_firewall_1 is Adding a Rule.
test_ids_1 is Processing Packet: "Hey I'm packet 2".
test_ids_2 is Processing Packet: "Hey I'm packet 2".
test_firewall_1 is Processing Packet: "Hey I'm packet 3".
test_firewall_1 is Adding a Rule.
test_ids_1 is Processing Packet: "Hey I'm packet 3".
test_ids_2 is Processing Packet: "Hey I'm packet 3".
PASS
Process exiting with code: 0
```
