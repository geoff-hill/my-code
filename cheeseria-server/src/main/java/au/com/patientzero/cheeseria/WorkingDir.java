package au.com.patientzero.cheeseria;

import java.io.File;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

/**
 * Cheeseria stores working information including content and sometimes data in a {@code .cheeseria} directory created in the User's home directory.
 * This class contains static helper methods to retrieve and create these files.
 * 
 * @
 */
public class WorkingDir {
  private static final Logger logger = LoggerFactory.getLogger(WorkingDir.class);
  private static final Path WORKING_DIR_PATH = Path.of(System.getProperty("user.home")).resolve(".cheeseria");

  /**
   * Gets the requested File from the working directory. If a file with that name already exists,
   * it is returned untouched.  If not, a new file is created, and the content of the given defaultContentClassPathResource
   * is copied to the file.
   * 
   * @param fileName The fileName in the working directory.
   * @param defaultContentClassPathResource The path to the default content, relative to the classpath
   * @return The file found in the working directory
   * @throws IOException
   */
  public static File getOrCreateFile(String fileName, String defaultContentClassPathResource) throws IOException {
    logger.debug("getOrCreateFile(\"{}\", \"{}\")", fileName, defaultContentClassPathResource);

    Files.createDirectories(WORKING_DIR_PATH); // throws only if cheeseria exists but is not a directory

    Path filePath = WORKING_DIR_PATH.resolve(fileName).normalize();
    
    if (Files.notExists(filePath)) {
      copyDefaultContentToNewFile(filePath, defaultContentClassPathResource);
    } else {
      logger.debug("Existing config found");
    }

    return filePath.toFile();
  }

  private static void copyDefaultContentToNewFile(Path  filePath, String defaultContentClassPathResource) throws IOException {
    try (var defaultContentStream = ClassLoader.getSystemResourceAsStream(defaultContentClassPathResource)) {
      if (defaultContentStream == null) {
        throw new DefaultContentNotFoundException("Required Classpath Resource: " + defaultContentClassPathResource);
      }
      logger.debug("Creating file with default content: " + filePath);
      Files.copy(defaultContentStream, filePath);
      logger.debug("Created successfully");
    }
  }
}

class DefaultContentNotFoundException extends IOException {
  DefaultContentNotFoundException(String msg) {
    super(msg);
  }
}
