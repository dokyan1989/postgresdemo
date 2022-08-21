package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dokyan1989/postgresdemo/helper/random"
)

var OrderIDs = []string{
	"8dcb3eff-7a20-4165-a96f-4e7185dba675",
	"0688ff9f-ff67-47f4-b607-003886c28c87",
	"c64f6fae-9a7e-4102-b4b4-693ea1b16004",
	"e8ecd0dc-f346-4002-9fc8-3b9815e05f6d",
	"4edb4b08-5df8-45d3-863d-89ac2498ea2a",
	"dc532165-98a8-4d3b-ab95-b4e7d274c4e0",
	"2d820c78-3f33-4cc5-906f-3cdd628e8d7e",
	"eb1818a9-cff0-43ad-9103-d4b8a38d7fac",
	"9fca1ddc-7d3c-47fe-b6be-382707e5ce08",
	"e8061dd7-e199-41b1-bb36-364941368dda",
	"ee65ca50-ace8-489c-83e4-b1ef8645e74b",
	"85e917da-ce61-493e-866a-984c3998f8b8",
	"232a89c5-f4b4-41a5-9e89-b1abc2518ef3",
	"faa7a8b7-dfa7-47c6-91e2-6cb4efa38a1e",
	"14a58825-b207-4e16-aba2-8c977148f2c3",
	"2837ed98-98b5-4105-8b06-4e18189eee05",
	"4b75b681-4b2f-4811-8a27-aecbf473221c",
	"8f335bec-55b7-41d2-8885-d17117b961cc",
	"cdc95169-fb5b-4153-a8d5-0c3b50b73349",
	"5b23f63a-49c6-4edf-863a-564dfc9267f6",
	"f399395a-9e3e-4b14-a993-41eb8a04e60c",
	"b53bc589-13a7-4ffa-a3ac-907918b5f3ce",
	"107fb224-8a8a-4cdb-bc64-98a87f96a421",
	"8cbb1bca-bdef-426b-a4da-9aa8f3d8d13e",
	"284e2131-4677-40e0-b3e7-001c09d9b34e",
	"469826a2-1b77-4170-9f33-605d225602c4",
	"81097bf9-c3bd-4d49-a3e3-49ade06c4333",
	"b916b91f-5fe3-43e4-bf29-3e5f5c9e58e0",
	"ab2030aa-9833-4005-8753-2b0ecfe28224",
	"83b1bf9f-d129-4b85-9fab-65b5c79818b8",
	"6bced5b0-ce21-46bb-934f-a29ceb046bbb",
	"71b96fb3-08f3-49ce-afb3-8134bf1d7f82",
	"5b015785-e3bc-47f6-890c-3d74c63f4ebd",
	"478bc450-f6e8-4b89-bc01-9931091c2e6c",
	"581aa2f0-8700-48e5-bcb0-ccc17e6ef2f8",
	"a3dec9a1-7ec1-41c7-a318-363f0c160017",
	"9225a6ff-06e4-4482-acd2-b55e65d9874c",
	"81fe21bd-f232-42d1-9a18-90306fc3db90",
	"3c36de4d-e3d3-46f9-9741-8a5566a9b81b",
	"ef8d7d35-38dc-4737-a8e9-d117d9cc6196",
	"00a84a46-4b84-445e-92bd-1db4492bf527",
	"9db26025-23c0-4652-83aa-efe84fce0c43",
	"1b0aa58b-0b0c-4d8c-b5f0-ad1dda7e032f",
	"767d41c6-5213-4e58-aae1-d45bec638e40",
	"72d79e5f-76e8-426c-b143-59898cdef694",
	"b98100b2-0920-436b-8c69-f2c87956c2da",
	"96f695b2-d44b-485d-8f43-cdad86f94486",
	"6de41798-5c5c-4c95-8564-89dab1bf9414",
	"f21c62ff-fcac-4874-8990-b4ffd3c97e94",
	"f9a903e0-2762-4497-96ab-d20fac381dbc",
	"167ed012-8168-4d7d-97f1-55e10876fa99",
	"de7baefb-91e0-4ae4-9c0b-f6e6d225ed28",
	"b839aa68-0ae5-495d-b9a2-872bd881af84",
	"e07a64d2-a496-4939-b3f5-9e724a77b5c8",
	"939ae7cb-cf94-4024-be78-cad9a8f9bbb1",
	"ef1faeaa-e9cb-4adf-bbe1-4855830d09bd",
	"f40ea6e7-b044-4730-bd14-63d49ef51bac",
	"daf9df7e-556c-44fc-ab08-2f1c4a6747d0",
	"d8f9e71f-a351-4591-82ce-a3a8f3294d9a",
	"2f07e7b8-42c3-4f5f-bbee-588956dc7e73",
	"b245920f-5a8e-4f73-8823-e9e854278910",
	"45fb5dcd-d6be-4627-9821-e7c627a34764",
	"f4b807d4-c610-4f25-953b-b1950ba1fdd4",
	"37db1c71-420e-467f-8692-1ed3a7577b22",
	"59c8b638-2ee1-4467-9326-ec665a444ccb",
	"eb8102a0-75b1-4027-8468-2d89c8df56e6",
	"2b2e18b8-b710-4693-bc72-87452380b5aa",
	"17d9be6d-1c28-4eda-a9a2-15351f4536d0",
	"b742a269-9a44-46c0-aa5a-90ed902ca03f",
	"936bdc42-9e3c-4a05-997a-adb108770fe6",
	"793f479e-a269-403e-85a7-4c4eedc66bc5",
	"92edc8a9-524a-4eb3-be45-5fca9347b765",
	"c27d2356-b534-4f39-99ed-46f911fa4dd6",
	"b137ffd2-c055-4902-965f-c4ce6710b0cc",
	"58a6167a-3f67-41e3-86d7-1b5bf2723097",
	"15eb875d-5346-4c66-a2db-1749c7fe99e3",
	"055758ce-a24e-40c6-9a17-f5acdfc45227",
	"10c6d36a-8194-4ccd-a545-4b91007bf0b9",
	"6c382adb-79c7-4341-9f66-ee982bfd4302",
	"8a8be429-09e8-486a-a07e-43f972d15ec4",
	"198c45fa-72cc-4035-93fb-2de7a2c14886",
	"8a549da2-275a-47ae-ae63-ce4cd9893e90",
	"82beb45f-46e1-4674-9908-466c6be77403",
	"97911b24-2192-4acb-957d-71a71e7e1688",
	"d7d54fc5-b82a-4bc8-8cb9-f0ae422f267f",
	"3e301837-a266-4492-8f13-2ff48475391a",
	"b7b68b0c-3fdd-40ad-a5a3-bb78b8d8668c",
	"28365606-d316-48e4-91e6-d60af6e4f816",
	"27305fe2-d017-4b62-8b26-1d69411bd696",
	"2f08a4dc-1e9e-4862-8fd5-caf7a79a45a0",
	"cc90a52d-96a0-4b54-b50d-1942a2ad6da0",
	"bd2521e1-58f5-42f3-881f-b145c12c3c56",
	"bde576c9-1745-4016-ae9e-7e1eef819d7d",
	"99cac383-ef75-4e7b-a3e6-879d9d63b009",
	"731547a0-f828-4b2d-9d37-0657093eeefb",
	"c80871f4-df9d-4b48-954a-0c13fbfa3e78",
	"e19eb89b-879b-4fc1-a833-a828cdd9d7b8",
	"6dd62809-e0c9-41b1-b779-f32b9cce57c9",
	"8b4acadb-dc62-4df9-9d88-473f8234ab1b",
	"39e10d01-e8c6-4eb4-a74a-acede7608ff1",
}

func main() {
	for {
		randomId := OrderIDs[random.Int(0, 99)]
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:3000/orders/%s", randomId), nil)
		if err != nil {
			panic(err)
		}

		client := http.DefaultClient
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Make an HTTP Request with Method:%s - URL:%s at %s\n", req.Method, req.URL, time.Now().Format(time.RFC3339))
		// body, err := ioutil.ReadAll(resp.Body)
		// if err != nil {
		// 	panic(err)
		// }
		fmt.Printf("Received an HTTP Response with StatusCode:%d - at %s\n", resp.StatusCode, time.Now().Format(time.RFC3339))

		delay := random.Int(0, 1000)
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
}

func PrettyStruct(data interface{}) (string, error) {
	val, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(val), nil
}
