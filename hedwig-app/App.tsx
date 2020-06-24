import React from 'react';
import { StyleSheet, Text, View, Platform, Button, Share } from 'react-native';
import { Notifications, AppLoading } from 'expo';
import * as Permissions from 'expo-permissions';
import Constants from 'expo-constants';
import * as Font from 'expo-font';
import { NewsCycle_400Regular } from '@expo-google-fonts/news-cycle'
import { SpaceMono_400Regular } from '@expo-google-fonts/space-mono'
import * as WebBrowser from 'expo-web-browser';

interface AppState {
	expoPushToken?: string
	fontsLoaded: boolean
}

export default class App extends React.Component<{}, AppState> {
	constructor(props: {}) {
		super(props)

		this.state = {
			fontsLoaded: false
		}
	}

	async registerForPushNotifications() {
		if (Constants.isDevice) {
			let { status } = await Permissions.getAsync(Permissions.NOTIFICATIONS)
			if (status != "granted") {
				const { status: statusResp } = await Permissions.askAsync(Permissions.NOTIFICATIONS)
				status = statusResp;
			}

			if (status != "granted") {
				alert("You must grant notification permissions in settings.")
			}
			let token = await Notifications.getExpoPushTokenAsync();
			console.log(token)
			this.setState({ expoPushToken: token });
		} else {
			alert("You must use a physical device to receive notifications.")
		}

		if (Platform.OS === 'android') {
			Notifications.createChannelAndroidAsync('default', {
				name: 'default',
				sound: true,
				priority: 'max',
				vibrate: [0, 250, 250, 250],
			});
		}
	}

	async loadFonts() {
		await Font.loadAsync({ SpaceMono_400Regular, NewsCycle_400Regular })
		this.setState({
			fontsLoaded: true
		})
	}

	componentDidMount() {
		this.registerForPushNotifications()
		this.loadFonts()
	}

	render() {
		if (!this.state.fontsLoaded) {
			return <AppLoading />
		}
		return (
			<View style={styles.container}>
				<Text style={[styles.title, styles.bottomPadded]}>Hedwig</Text>
				<Text style={styles.bottomPadded}>Visit https://willbarkoff.dev/hedwig for docs.</Text>
				{this.state.expoPushToken && <View>
					<Text style={[styles.bottomPadded, styles.monospaced]} selectable={true}>{this.state.expoPushToken}</Text>
					<Button title="Share" onPress={() => Share.share({ message: this.state.expoPushToken! })} />
					<Button title="Documentation" onPress={() => WebBrowser.openBrowserAsync("https://willbarkoff.dev/hedwig")} />
				</View>}
			</View>
		);
	}
}

const styles = StyleSheet.create({
	title: {
		fontSize: 36,
		fontFamily: "NewsCycle_400Regular",
	},
	monospaced: {
		fontFamily: "SpaceMono_400Regular",
	},
	bottomPadded: {
		paddingBottom: 10,
	},
	container: {
		flex: 1,
		backgroundColor: '#fff',
		alignItems: 'center',
		justifyContent: 'center',
	},
});
