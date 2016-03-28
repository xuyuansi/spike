package spike.xsi.encrypt;

import org.apache.commons.codec.binary.Base64;

public class App {
    public static void main(String[] args) {
        String clientId = "9PgZxZ78iKd9Io6G2f0OGw";
        String clientSecret = "CD0F0668-4B7E-44CD-AF39-43C32937F380";

        String base64 = Base64.encodeBase64String((clientId + ":" + clientSecret).getBytes());
        System.out.println("Encoded : " + base64);

        byte[] bytes = Base64.decodeBase64(base64);
        String decoded = new String(bytes);

        System.out.println("Decoded : " + decoded);
    }
}
