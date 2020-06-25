---
layout: page
title: notify
---

<script>
	(async () => {
	let resp = await fetch("https://exp.host/--/api/v2/push/send", {
		"credentials": "omit",
		"body": "[{\"to\":\"ExponentPushToken[kDrrfaC-vP1_HBU69ZThvr]\",\"title\":\"Title\",\"body\":\"Body\",\"ttl\":10}]",
		"method": "POST",
		"mode": "cors"
	});
	console.log(resp)
})();
</script>