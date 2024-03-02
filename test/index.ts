import "dotenv/config"
import { Client, GatewayIntentBits } from 'discord.js';

const client = new Client({ intents: [GatewayIntentBits.Guilds] });
if (client == null) {
    console.log(`Client Not Found!`)
}

client.on('ready', () => {
    console.log(`Logged in as ${client?.user?.tag}!`);
});

client.on('interactionCreate', async interaction => {
    if (!interaction.isChatInputCommand()) return;

    if (interaction.commandName === 'ping') {
        await interaction.reply('Pong!');
    }
});

// client.login(Deno.env.get('TOKEN'));
while (true) {}
