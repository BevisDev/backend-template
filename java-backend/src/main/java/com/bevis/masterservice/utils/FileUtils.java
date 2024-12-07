package com.bevis.masterservice.utils;

import java.io.File;
import java.io.IOException;
import java.io.InputStream;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.nio.file.StandardCopyOption;

import org.apache.commons.lang3.StringUtils;

public class FileUtils {

    /**
     * path replaced to matches operating system
     * Window: (\); Linux (/)
     *
     * @param pathName path name
     * @return pathName
     */
    public static String cleanPath(String pathName) {
        return StringUtils.replace(pathName, "%", File.separator);
    }

    public static String getExtension(String fileName) {
        int index = StringUtils.lastIndexOf(fileName, ".");
        return StringUtils.substring(fileName, index + 1);
    }

    public static void createDirectories(String path) throws IOException {
        Path dir = Paths.get(path);
        if (!Files.exists(dir)) {
            Files.createDirectories(dir);
        }
    }

    /**
     * upload file
     *
     * @param inputStream
     * @param desFile
     * @throws IOException
     */
    public static void uploadFile(InputStream inputStream, Path desFile) throws IOException {
        Files.copy(inputStream, desFile, StandardCopyOption.REPLACE_EXISTING);
    }
}
